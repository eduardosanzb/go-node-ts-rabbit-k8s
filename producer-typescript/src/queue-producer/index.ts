var q = "tasks";
import client, { Connection, Channel } from "amqplib";

interface QueueProducer {
  send: (message: string) => Promise<void>;
  isReady: boolean;
}

export class RabbitMQProducer implements QueueProducer {
  isReady = false;
  channel: Channel | undefined;
  connection: Connection | undefined;
  _queue_name: string;
  _BROKER_URL: string;

  constructor() {
    this._queue_name = process.env.CHANNEL_NAME || "myQueue";
    this._BROKER_URL =
      process.env.BROKER_URL ||
      "amqp://username:password@host.docker.internal:5672";
    this._init();
  }

  async _init() {
    try {
      this.connection = await client.connect(this._BROKER_URL, "heartbeat=30");
      const channel: Channel = await this.connection.createChannel();
      await channel.assertQueue(this._queue_name);
      this.channel = channel;

      this.isReady = true;
    } catch (e) {
      console.error(e);
      console.trace("====");
      process.exit(1);
    }
  }

  async send(message: string) {
    await this.channel?.sendToQueue(this._queue_name, Buffer.from(message));
  }
}

export function createQueueProducer(): QueueProducer {
  // maybe use DI or something else when fancy
  const queue = new RabbitMQProducer();
  process.stdout.write("Connecting to the queue ..." + "\r");
  return queue;
}
