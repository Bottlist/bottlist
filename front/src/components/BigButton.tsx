import { Link as RouterLink } from 'react-router-dom';

export const BigButton = (props: { children: string; link: string }) => (
  <RouterLink
    to={props.link}
    style={{
      width: 280,
      height: 60,
      borderRadius: 30,
      boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25)',
      backgroundColor: 'white',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      textDecoration: 'none',
      letterSpacing: 20,
      fontSize: 32,
      color: 'black',
    }}
  >
    {props.children}
  </RouterLink>
);
