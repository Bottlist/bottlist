import { FormControl, FormHelperText, Typography } from '@mui/material';
import { Input } from '@mui/base/Input';
import { FieldValues, Path, useFormContext } from 'react-hook-form';

export const GrayTextField = <T extends FieldValues>(props: {
  _key: Path<T>;
  helperText?: string;
  label: string;
  alignLabel?: 'left' | 'center';
}) => {
  const {
    register,
    formState: { errors },
  } = useFormContext<T>();
  const { _key } = props;
  return (
    <FormControl error={!!errors[_key]}>
      <label>
        <Typography fontSize={13} textAlign={props.alignLabel ?? 'left'}>
          {props.label}
        </Typography>
      </label>
      <Input
        {...register(_key)}
        slotProps={{
          input: {
            style: {
              width: 210,
              height: 40,
              backgroundColor: '#D9D9D9',
              boxShadow: '0px 2px 4px 0px rgba(0, 0, 0, 0.25) inset',
            },
          },
        }}
      />
      <FormHelperText>
        {errors[_key]?.message?.toString() ?? props.helperText}
      </FormHelperText>
    </FormControl>
  );
};
