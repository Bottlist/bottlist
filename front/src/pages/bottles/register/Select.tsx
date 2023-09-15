import {
  Box,
  InputAdornment,
  Paper,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import PlaceIcon from '@mui/icons-material/Place';
import GoogleMapReact from 'google-map-react';
import { useQuery } from '@tanstack/react-query';
import { request } from '../../../utils/axiosUtils';
import { useFormContext } from 'react-hook-form';
import { FormType } from '.';

const key = process.env.REACT_APP_GOOGLE_MAP_KEY ?? '';

export const Select = () => {
  const { data } = useQuery({
    queryKey: ['shops'],
    queryFn: () =>
      request({
        url: '/shops',
        method: 'get',
      }).then((r) => r.data.shops),
  });

  const { setValue } = useFormContext<FormType>();
  return (
    <Stack>
      <TextField
        label="お店の名前で検索"
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <SearchIcon />
            </InputAdornment>
          ),
        }}
      />
      <Stack direction="row">
        <PlaceIcon />
        <Typography>現在地から探す</Typography>
      </Stack>
      <Box sx={{ height: '50vh', width: '100%' }}>
        <GoogleMapReact
          bootstrapURLKeys={{ key }}
          defaultCenter={{
            lat: 34.75,
            lng: 137.75,
          }}
          defaultZoom={11}
        ></GoogleMapReact>
      </Box>
      <Paper>
        {data?.map((shop) => (
          <Typography key={shop.id} onClick={() => setValue('shop', shop.id)}>
            {shop.name} {shop.address}
          </Typography>
        ))}
      </Paper>
    </Stack>
  );
};
