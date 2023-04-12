import {
  Card,
  Grid,
  CardMedia,
  CardContent,
  Stack,
  Divider,
  CardActions,
  Typography,
} from '@mui/material';
import { ReactNode } from 'react';

export const BottleCard = (props: {
  upperText: string;
  lowerText: string;
  children: ReactNode;
}) => (
  <Card>
    <Grid container alignItems="center">
      <Grid item xs={2}>
        <CardMedia
          component="img"
          image="https://www.yomeishu-online.jp/yomeishu-online_wp/wp-content/uploads/2020/08/herb_set700x300_p01.jpg"
        />
      </Grid>
      <Grid item xs={7}>
        <CardContent>
          <Stack direction="row" spacing={2}>
            <Typography>{props.upperText}</Typography>
          </Stack>
          <Divider />
          <Stack direction="row" spacing={2}>
            <Typography>{props.lowerText}</Typography>
          </Stack>
        </CardContent>
      </Grid>
      <Grid item xs={3} flexGrow={0}>
        <CardActions>{props.children}</CardActions>
      </Grid>
    </Grid>
  </Card>
);
