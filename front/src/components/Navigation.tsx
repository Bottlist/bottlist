import { BottomNavigation, BottomNavigationAction } from '@mui/material';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import SanitizerRoundedIcon from '@mui/icons-material/SanitizerRounded';
import AccountBoxRoundedIcon from '@mui/icons-material/AccountBoxRounded';
import { Link, LinkProps } from 'react-router-dom';

/** react-router-dom's Link component doesn't work well with BottomNavigation by default, create thin wrapper here */
const MyLink = (props: { href: string } & Omit<LinkProps, 'to'>) => (
  <Link to={props.href} {...props} />
);

export const Navigation = () => (
  <BottomNavigation
    showLabels
    sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }}
  >
    <BottomNavigationAction
      icon={<HomeRoundedIcon />}
      label="ホーム"
      LinkComponent={MyLink}
      href="/top"
    />
    <BottomNavigationAction
      label="ボトル追加"
      icon={<SanitizerRoundedIcon />}
      LinkComponent={MyLink}
      href="/bottle/register"
    />
    <BottomNavigationAction
      label="アカウント"
      icon={<AccountBoxRoundedIcon />}
      LinkComponent={MyLink}
      href="/"
    />
  </BottomNavigation>
);
