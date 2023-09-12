import { RouteObject } from 'react-router-dom';
import { Register } from './register';
import { Update } from './Update';
import { Select } from './register/Select';
import { Form } from './register/Form';

export const bottleRouter: RouteObject[] = [
  { path: 'update', element: <Update /> },
  {
    path: 'register',
    element: <Register />,
    children: [
      { path: '', element: <Form /> },
      { path: 'select-shop', element: <Select /> },
    ],
  },
];
