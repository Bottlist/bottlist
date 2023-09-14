import { BottomNavigation, BottomNavigationAction } from '@mui/material';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import SwitchAccountIcon from '@mui/icons-material/SwitchAccount';
import SettingsIcon from '@mui/icons-material/Settings';
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
      href="/shops"
    />
    <BottomNavigationAction
      label="顧客様一覧"
      icon={<SwitchAccountIcon />}
      LinkComponent={MyLink}
      href="/shops/users"
    />
    <BottomNavigationAction
      label="設定"
      icon={<SettingsIcon />}
      LinkComponent={MyLink}
      href="/users/profile"
    />
  </BottomNavigation>
);
