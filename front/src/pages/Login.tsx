import {
  Box,
  Button,
  Container,
  Paper,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import { VerticalCenter } from '../components/VerticalCenter';

export const Login = () => {
  return (
    <VerticalCenter>
      <Stack spacing={5}>
        <Typography variant="h4">BOTTLIST</Typography>
        <Paper>
          <Container>
            <Stack spacing={5} marginY={5} alignItems="center">
              <TextField placeholder="username or E-mail Address" />
              <TextField placeholder="Password" />
              <Box>
                <Button>ログイン</Button>
              </Box>
            </Stack>
          </Container>
        </Paper>
        <Box>
          <Button color="secondary">Lost your password?</Button>
        </Box>
      </Stack>
    </VerticalCenter>
  );
};
