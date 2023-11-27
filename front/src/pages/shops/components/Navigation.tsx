import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import SwitchAccountIcon from '@mui/icons-material/SwitchAccount';
import SettingsIcon from '@mui/icons-material/Settings';
import { Navigation as Base } from '../../../components/Navigation';

const NAVIGATIONS = [
  {
    label: 'ホーム',
    icon: <HomeRoundedIcon fontSize="large" />,
    href: '/shops',
  },
  {
    label: 'ボトル追加',
    icon: <SwitchAccountIcon fontSize="large" />,
    href: '/shops/users',
  },
  {
    label: 'アカウント',
    icon: <SettingsIcon fontSize="large" />,
    href: '/shops/settings',
  },
];

export const Navigation = () => <Base navigations={NAVIGATIONS} />;
