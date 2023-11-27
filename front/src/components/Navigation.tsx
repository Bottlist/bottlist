import { BottomNavigation, BottomNavigationAction } from '@mui/material';
import { Link, LinkProps, useLocation } from 'react-router-dom';

/** react-router-dom's Link component doesn't work well with BottomNavigation by default, create thin wrapper here */
const MyLink = (props: { href: string } & Omit<LinkProps, 'to'>) => (
  <Link to={props.href} {...props} />
);

type Props = {
  navigations: { label: string; icon: JSX.Element; href: string }[];
};

export const Navigation = (props: Props) => {
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
      {props.navigations.map((navigation) => (
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
