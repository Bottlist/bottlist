import { Profile } from './Profile';
import { Register } from './Register';

export const usersRouter = [
  {
    path: 'register',
    element: <Register />,
  },
  { path: 'profile', element: <Profile /> },
];
