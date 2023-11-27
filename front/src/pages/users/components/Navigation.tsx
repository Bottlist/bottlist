import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import SanitizerRoundedIcon from '@mui/icons-material/SanitizerRounded';
import AccountBoxRoundedIcon from '@mui/icons-material/AccountBoxRounded';
import { Navigation as Base } from '../../../components/Navigation';

const NAVIGATIONS = [
  { label: 'ホーム', icon: <HomeRoundedIcon fontSize="large" />, href: '/top' },
  {
    label: 'ボトル追加',
    icon: <SanitizerRoundedIcon fontSize="large" />,
    href: '/bottle/register',
  },
  {
    label: 'アカウント',
    icon: <AccountBoxRoundedIcon fontSize="large" />,
    href: '/users/profile',
  },
];

export const Navigation = () => <Base navigations={NAVIGATIONS} />;
