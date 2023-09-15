import { PickLogin } from './PickLogin';
import { Profile } from './Profile';
import { Register } from './Register';

export const usersRouter = [
  {
    path: 'pick-login',
    element: <PickLogin />,
  },
  {
    path: 'register',
    element: <Register />,
  },
  { path: 'profile', element: <Profile /> },
];
