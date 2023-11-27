import { Login } from './Login';
import { PickLogin } from './PickLogin';
import { Profile } from './Profile';
import { Register } from './Register';
import { Top } from './Top';

export const usersRouter = [
  {
    path: '',
    element: <Top />,
  },
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
  { path: 'profile', element: <Profile /> },
];
