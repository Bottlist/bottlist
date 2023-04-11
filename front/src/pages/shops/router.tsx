import { Register } from './Register';
import { Top } from './Top';

export const shopsRouter = [
  {
    path: 'register',
    element: <Register />,
  },
  {
    path: '',
    element: <Top />,
  },
];
