import { Box } from '@mui/material';
import { Logo } from './Logo';
import { BigButton } from './BigButton';

export function ChoicesLayout(props: {
  items: {
    text: string;
    link: string;
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
          <BigButton link={item.link}>{item.text}</BigButton>
        </Box>
      ))}
    </>
  );
}
