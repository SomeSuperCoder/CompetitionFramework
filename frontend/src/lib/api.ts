import axios from 'axios';

const RPC_URL = 'http://localhost:8899/rpc';

interface JSONRPCResponse<T> {
  result: T;
  error: string;
  id: number;
}

export async function invokeJSONRPC<T>(
  functionName: string,
  data: object
): Promise<JSONRPCResponse<T>> {
  const serializedObject = JSON.stringify(data);
  const requestBody = `
{
    "jsonrpc": "2.0",
    "method": "${functionName}",
    "params": [${serializedObject}],
    "id": 1
}`;
  return (
    await axios.post<JSONRPCResponse<T>>(RPC_URL, requestBody, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
  ).data;
}

export async function getCompetitons(): Promise<JSONRPCResponse<unknown>> {
  return invokeJSONRPC('Competition.FindAll', {});
}
