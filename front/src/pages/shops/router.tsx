import { Register } from './Register';
import { Top } from './Top';
import { UserList } from './UserList';

export const shopsRouter = [
  {
    path: 'register',
    element: <Register />,
  },
  {
    path: '',
    element: <Top />,
  },
  {
    path: 'users',
    element: <UserList />,
  },
];
