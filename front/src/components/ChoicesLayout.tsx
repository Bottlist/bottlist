import { Box, Button, Typography } from '@mui/material';
import { Link as RouterLink } from 'react-router-dom';
import { Logo } from './Logo';

export function ChoicesLayout(props: {
  items: {
    text: string;
    link?: string;
  }[];
}) {
  return (
    <>
      <Box
        sx={{
          position: 'relative',
          top: 165,
          marginX: 'auto',
          width: 'fit-content',
        }}
      >
        <Logo />
      </Box>
      {props.items.map((item, i) => (
        <Box
          key={i}
          sx={{
            position: 'relative',
            top: 438 + (55 + 60) * i,
            left: 57,
            width: 'fit-content',
          }}
        >
          <PickButton link={item.link}>{item.text}</PickButton>
        </Box>
      ))}
    </>
  );
}

const PickButton = (props: { children: string; link?: string }) => {
  const { children, link } = props;
  return (
    <Button
      color="secondary"
      fullWidth
      sx={{ padding: 2, borderRadius: 30, height: 60, width: 280 }}
      {...(link && { component: RouterLink, to: link })}
    >
      <Typography variant="h4" letterSpacing={20}>
        {children}
      </Typography>
    </Button>
  );
};
