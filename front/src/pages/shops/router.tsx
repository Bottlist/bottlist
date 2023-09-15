import { PickLogin } from './PickLogin';
import { Register } from './Register';
import { Top } from './Top';
import { Index } from './users/Index';
import { View } from './users/View';

export const shopsRouter = [
  {
    path: 'pick-login',
    element: <PickLogin />,
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
];
