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

mock.onPost('/auth/login/user').reply(200);
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
mock.onGet('/me').reply<ResponseData<'/me', 'get'>>(200, {
  first_name: 'string',
  first_name_huri: 'string',
  last_name: 'string',
  last_name_huri: 'string',
  screen_name: 'string',
  birthday: '2019-08-24',
  email: 'user@example.com',
  img: 'http://example.com',
});
mock
  .onPost('/auth/password/reset/link')
  .reply<ResponseData<'/auth/password/reset/link', 'post'>>(200);
mock.onGet('/shops/bottles').reply<ResponseData<'/shops/bottles', 'get'>>(200, {
  bottles: [
    {
      type: 'whisky',
      name: 'string',
      expires_at: '2019-08-24',
      amount: 0.4,
      status: 'pending',
      id: '1',
      username: 'user',
    },
    {
      type: 'whisky',
      name: 'string',
      expires_at: '2019-08-24',
      amount: 0.4,
      status: 'pending',
      id: '2',
      username: 'user',
    },
    {
      type: 'whisky',
      name: 'string',
      expires_at: '2019-08-24',
      amount: 0.4,
      status: 'approved',
      id: '3',
      username: 'user',
    },
  ],
});
mock.onGet('/users').reply<ResponseData<'/users', 'get'>>(200, {
  users: [
    { id: 'string', name: 'string', img: 'string', birthday: '2019-08-24' },
    { id: 'string2', name: 'string', img: 'string', birthday: '2019-08-24' },
  ],
});
mock.onGet('/users/1').reply<ResponseData<'/users/{id}', 'get'>>(200, {
  name: 'string',
  img: 'string',
  birthdate: '2019-08-24',
  bottles: [
    {
      amount: 0.4,
      name: 'string',
      shop_name: 'string',
      type: 'whisky',
      expires_at: '2019-08-24',
    },
  ],
});
