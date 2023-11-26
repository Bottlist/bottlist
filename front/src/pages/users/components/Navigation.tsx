import { BottomNavigation, BottomNavigationAction } from '@mui/material';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import SanitizerRoundedIcon from '@mui/icons-material/SanitizerRounded';
import AccountBoxRoundedIcon from '@mui/icons-material/AccountBoxRounded';
import { Link, LinkProps, useLocation } from 'react-router-dom';

/** react-router-dom's Link component doesn't work well with BottomNavigation by default, create thin wrapper here */
const MyLink = (props: { href: string } & Omit<LinkProps, 'to'>) => (
  <Link to={props.href} {...props} />
);

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

export const Navigation = () => {
  const { pathname } = useLocation();
  return (
    <BottomNavigation
      value={pathname}
      showLabels
      sx={{
        position: 'fixed',
        bottom: 0,
        left: 0,
        right: 0,
        height: '100px',
        borderRadius: '20px 20px 0px 0px',
      }}
    >
      {NAVIGATIONS.map((navigation) => (
        <BottomNavigationAction
          key={navigation.href}
          icon={navigation.icon}
          label={navigation.label}
          LinkComponent={MyLink}
          href={navigation.href}
          value={navigation.href}
        />
      ))}
    </BottomNavigation>
  );
};
