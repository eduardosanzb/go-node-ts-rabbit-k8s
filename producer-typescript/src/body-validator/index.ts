export function validIpV4(ip:string){
  if (!ip){
    return false
  }

  const parts = ip.split('.');

  if(parts.length !== 4){
    return false
  }

  for(let p of parts){
    // little trick given that our lovely JS will ignore spaces when parseInt
    p.replaceAll(' ', 'x')

    if(
      (p.length === 0) ||
      isNaN(parseInt(Number(p) as unknown as string)) ||
      (parseInt(p) < 0 || parseInt(p) > 255) || // the previous rule ensure that we have digits
      (p.length > 1 && p.at(0) === '0')
    ){
      return false;
    }
  }

  return true;
}

export function validTimestamp(time:string | number){
  if(!time){
    return false
  }
  return new Date(Number(time)).getTime() > 0;
}


function isObject(o:any):any{
  return !Array.isArray(o) ||
    (typeof o !== 'object') ||
    (o === null)
}

export function validMessage(data:any): data is {message:any} {
  if(!isObject(data)){
    return false
  }
  return Object.keys(data?.message??{}).length >= 1 ;
}

export function containsSender(data:any): data is{sender:string}{
  if(!isObject(data)){
    return false
  }

  return typeof data?.sender === 'string' ?? false;
}

export function onlyValidKeys(data: any){
  if(!isObject(data)){
    return false
  }

  const validKeys = {'ts': true,'sender': true,'message': true,'sent-from-ip': true,'priority': true}
  let count = 0;
  for(const k of Object.keys(data)){
    if(!k as unknown as string in validKeys){
      return false
    }
    count++;
  }

  return count === Object.keys(validKeys).length;
}

export default function bodyValidator(data:any){
  return (
    validIpV4(data['sent-from-ip']) &&
    validTimestamp(data['ts']) &&
    containsSender(data) &&
    validMessage(data) &&
    onlyValidKeys(data)
  )
}
