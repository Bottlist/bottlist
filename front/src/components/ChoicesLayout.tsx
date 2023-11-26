import { Box, Stack } from '@mui/material';
import { BigButton } from './button/BigButton';
import { CenterLogo } from './logo/CenterLogo';

export function ChoicesLayout(props: {
  items: {
    text: string;
    link: string;
  }[];
}) {
  return (
    <>
      <CenterLogo />
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
