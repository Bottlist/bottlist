import { yupResolver } from '@hookform/resolvers/yup';
import { useForm, FormProvider } from 'react-hook-form';
import * as yup from 'yup';
import { Navigation } from '../../users/components/Navigation';
import { DatePicker } from '../../../components/DatePicker';
import { TextField } from '../../../components/TextField';
import {
  Autocomplete,
  Paper,
  Stack,
  ToggleButton,
  ToggleButtonGroup,
  Typography,
  TextField as MuiTextField,
} from '@mui/material';
import { useQuery } from '@tanstack/react-query';
import { request } from '../../../utils/axiosUtils';
import { UpperLeftLogo } from '../../../components/logo/UpperLeftLogo';
import { BOTTLE_TYPE } from '../../../constants';
import { Sake } from '../../../components/sake/Sake';
import { RejectedBottles } from '../../../components/RejectedBottles';
import { Button } from '../../../components/button/Button';

const schema = yup
  .object({
    date: yup.date().required(),
    shop: yup.string().required(),
    name: yup.string().required(),
    type: yup.string().oneOf(BOTTLE_TYPE.map((b) => b.key)),
  })
  .required();
export type FormType = yup.InferType<typeof schema>;

export const Register = () => {
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const submit = () => {
    console.log('TBD');
  };
  const { data } = useQuery({
    queryKey: ['shops'],
    queryFn: () =>
      request({
        url: '/shops',
        method: 'get',
      }).then((r) => r.data.shops),
  });

  return (
    <>
      <UpperLeftLogo />
      <Stack spacing={3}>
        <FormProvider {...methods}>
          <form onSubmit={methods.handleSubmit(submit)}>
            <Stack
              spacing={2}
              marginTop="77px"
              paddingY="10px"
              alignItems="center"
              width="250px"
              marginX="auto"
            >
              <Typography variant="h5" textAlign="center">
                種類選択
              </Typography>
              <Paper sx={{ width: 'fit-content' }}>
                <ToggleButtonGroup
                  color="primary"
                  exclusive
                  {...methods.register('type')}
                  sx={{ width: 'fit-content' }}
                >
                  {BOTTLE_TYPE.map((b) => (
                    <ToggleButton key={b.key} value={b.key}>
                      <Sake type={b.key} />
                    </ToggleButton>
                  ))}
                </ToggleButtonGroup>
              </Paper>
              <TextField placeholder="銘柄入力" _key="name" />
              <Autocomplete
                options={data?.map((s) => ({ id: s.id, label: s.name })) ?? []}
                fullWidth
                renderInput={(params) => (
                  <MuiTextField {...params} label="店検索" />
                )}
              />
              <DatePicker
                label="開封日入力"
                control={methods.control}
                _key="date"
              />
              <Button color="primary" type="submit" width="fit-content">
                追加
              </Button>
            </Stack>
          </form>
        </FormProvider>
        <RejectedBottles />
      </Stack>
      <Navigation />
    </>
  );
};
