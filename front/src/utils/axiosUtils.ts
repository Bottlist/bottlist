import axios, { AxiosResponse } from 'axios';
import * as schemaHelper from './schemaHelper';
import MockAdapter from 'axios-mock-adapter';
import { ResponseData } from './schemaHelper';

export type AxiosConfigWrapper<
  Path extends schemaHelper.UrlPaths,
  Method extends schemaHelper.HttpMethods,
> = {
  url: Path;
  method: Method & schemaHelper.HttpMethodsFilteredByPath<Path>;
  params?: schemaHelper.RequestParameters<Path, Method>;
  data?: schemaHelper.RequestData<Path, Method>;
};

const client = axios.create({ baseURL: process.env.REACT_APP_API_URL });

export function request<
  Path extends schemaHelper.UrlPaths,
  Method extends schemaHelper.HttpMethods,
>(config: AxiosConfigWrapper<Path, Method>) {
  return client.request<
    schemaHelper.ResponseData<Path, Method>,
    AxiosResponse<schemaHelper.ResponseData<Path, Method>>,
    AxiosConfigWrapper<Path, Method>['data']
  >(config);
}

const mock = new MockAdapter(client);

mock.onPost('/login').reply(200);
mock.onGet('/bottles').reply<ResponseData<'/bottles', 'get'>>(200, {
  bottles: [
    {
      type: 'whisky',
      name: 'string',
      expires_at: '2019-08-24',
      amount: 0.4,
      status: { status: 'approved', reason: 'string' },
      shop: { name: 'string', id: 'string' },
      id: '1',
    },
    {
      type: 'whisky',
      name: 'string',
      expires_at: '2019-08-24',
      amount: 0.4,
      status: { status: 'pending', reason: 'string' },
      shop: { name: 'string', id: 'string' },
      id: '2',
    },
    {
      type: 'whisky',
      name: 'string',
      expires_at: '2019-08-24',
      amount: 0.4,
      status: { status: 'rejected', reason: 'string' },
      shop: { name: 'string', id: 'string' },
      id: '3',
    },
  ],
});
mock.onGet('/shops').reply<ResponseData<'/shops', 'get'>>(200, {
  shops: [
    {
      lat: 34.75,
      lng: 137.75,
      name: 'BEBER',
      id: 'string',
      address: '浜松市中区本城町',
    },
  ],
});
