import http, {ServerResponse, Server } from 'node:http'

type typeOK = (args:Parameters<typeof OK>[1])=>void
type typeNotOK = (args:Parameters<typeof NotOK>[1])=>void

export function createServer(port: number, path:string, fn: (data:string, OK: typeOK, NotOK: typeNotOK) => any): Server {
  const server = http.createServer()

  server.on('request', (request, response) => {
    const { method, url } = request;
    if(method !== 'POST'){
      return NotOK(response, { message: 'This Service only accepts POST requests', statusCode: 405 })
    }

    if(url !== path){
      const message = `the path ${url} does not exist`;
      return NotOK(response, { message, statusCode: 404 })
    }

    const body:Buffer[] = [];
    request
      .on('data', body.push.bind(body))
      .on('end', async () => {
        const data = Buffer.concat(body).toString();
        return fn(data, OK.bind(null, response), NotOK.bind(null, response))
      })
      .on('error', (e) => NotOK(response, { message: e.message }))
  });

  server.listen(port);
  console.log(`Server up & running on port: ${port}`)

  return server;
}

function NotOK(
  response:ServerResponse,
  {message = 'something went wrong', statusCode = 500}:Partial<{message: string, statusCode:number}> = {}
){
  const headers = { 'Content-Type': 'text/plain'}
  response.writeHead(statusCode, headers);
  response.write(message);
  response.end();
}

function OK(
  response:ServerResponse,
  message:string | void  = ''
){
  let headers = { 'Content-Type': 'text/plain' }
  if(message?.at(0) === '{' && message?.at(-1) === '}'){
    headers['Content-Type'] = 'application/json'
  }

  response.writeHead(200, headers);
  response.write(message);
  response.end();
}
