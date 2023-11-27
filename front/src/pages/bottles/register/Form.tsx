import { useFormContext } from 'react-hook-form';
import { DatePicker } from '../../../components/DatePicker';
import { TextField } from '../../../components/TextField';
import {
  Box,
  Button,
  Container,
  Divider,
  MenuItem,
  Paper,
  Stack,
  Typography,
} from '@mui/material';
import { Link } from 'react-router-dom';
import { FormType } from '.';
import { useQuery } from '@tanstack/react-query';
import { request } from '../../../utils/axiosUtils';
import { BOTTLE_TYPE } from '../../../constants';

export const Form = () => {
  const { control } = useFormContext<FormType>();
  const { data } = useQuery({
    queryKey: ['bottles'],
    queryFn: () =>
      request({
        url: '/bottles',
        method: 'get',
      }).then((r) => r.data.bottles),
  });
  return (
    <Container>
      <Stack spacing={2}>
        <TextField label="種類選択" _key="type" select>
          {BOTTLE_TYPE.map((b) => (
            <MenuItem key={b.key} value={b.key}>
              {b.label}
            </MenuItem>
          ))}
        </TextField>
        <TextField placeholder="銘柄入力" _key="type" />
        {/* <TextField placeholder="店検索" _key="shop" disabled /> */}
        <Link to="select-shop">店検索</Link>
        <DatePicker label="開封日入力" control={control} _key="date" />
        <Box textAlign="center">
          <Button color="secondary" type="submit">
            追加
          </Button>
        </Box>
      </Stack>
      <Divider />
      <Typography variant="h6">申請中・差戻しボトル一覧</Typography>
      {data
        ?.filter((b) => b.status.status !== 'approved')
        .map((bottle) => (
          <Paper key={bottle.id}>
            <Typography>{bottle.shop.name}</Typography>
            <Typography>{bottle.status.status}</Typography>
            <Typography>{bottle.name}</Typography>
            <Typography>{bottle.expires_at}</Typography>
          </Paper>
        ))}
    </Container>
  );
};
