import axios, { AxiosResponse } from 'axios';
import * as schemaHelper from './schemaHelper';
import MockAdapter from 'axios-mock-adapter';

export type AxiosConfigWrapper<
  Path extends schemaHelper.UrlPaths,
  Method extends schemaHelper.HttpMethods
> = {
  url: Path;
  method: Method & schemaHelper.HttpMethodsFilteredByPath<Path>;
  params?: schemaHelper.RequestParameters<Path, Method>;
  data?: schemaHelper.RequestData<Path, Method>;
};

const client = axios.create({ baseURL: process.env.REACT_APP_API_URL });

export function request<
  Path extends schemaHelper.UrlPaths,
  Method extends schemaHelper.HttpMethods
>(config: AxiosConfigWrapper<Path, Method>) {
  return client.request<
    schemaHelper.ResponseData<Path, Method>,
    AxiosResponse<schemaHelper.ResponseData<Path, Method>>,
    AxiosConfigWrapper<Path, Method>['data']
  >(config);
}

const mock = new MockAdapter(client);

mock.onPost('/login').reply(200);
