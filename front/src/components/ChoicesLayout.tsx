import { Box, Stack } from '@mui/material';
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
      <Box position="absolute" top={165} width="100%" textAlign="center">
        <Logo />
      </Box>
      <Box display="grid" sx={{ placeItems: 'center' }} height="100vh">
        <Stack spacing={10} width="fit-content">
          {props.items.map((item, i) => (
            <BigButton key={i} link={item.link}>
              {item.text}
            </BigButton>
          ))}
        </Stack>
      </Box>
    </>
  );
}
