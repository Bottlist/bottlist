import { LinkProps, Link as RouterLink } from 'react-router-dom';

export const Link = (props: LinkProps) => (
  <RouterLink
    {...props}
    style={{ textDecoration: 'none', color: 'black', ...props.style }}
  />
);
