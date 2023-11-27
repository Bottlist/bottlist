import { PickLogin } from './PickLogin';
import { Register } from './Register';
import { Top } from './Top';
import { Index } from './users/Index';
import { View } from './users/View';
import { Settings } from './Settings';
import { Login } from './Login';

export const shopsRouter = [
  {
    path: 'pick-login',
    element: <PickLogin />,
  },
  {
    path: 'login',
    element: <Login />,
  },
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
    element: <Index />,
  },
  {
    path: 'users/:id',
    element: <View />,
  },
  {
    path: 'settings',
    element: <Settings />,
  },
];
