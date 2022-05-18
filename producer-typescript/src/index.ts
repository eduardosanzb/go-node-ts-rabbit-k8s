import validateBody from "./body-validator";
import { createQueueProducer } from "./queue-producer";
import { createServer, ServerHandler } from "./server";

const {
  PATH_TO_INSERT_INTO_QUEUE = "/to-queue",
  PORT_SERVER = 8000
} =  process.env;

const queue = createQueueProducer();

const handler: ServerHandler = async (rawBody, OK, NotOK) => {
  const isValid = validateBody(JSON.parse(rawBody));

  if (!isValid) {
    return NotOK({
      message: "The request body is incorrect",
      statusCode: 400,
    });
  }

  await queue.send(rawBody);
  OK();
};

createServer(Number(PORT_SERVER), PATH_TO_INSERT_INTO_QUEUE, handler);
