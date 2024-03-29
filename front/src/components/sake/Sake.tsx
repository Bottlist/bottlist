import { components } from '../../schema';
import brandy from './brandy.png';
import shochu from './shochu.png';
import whisky from './Whiske.png';

type BottleType = components['schemas']['bottle']['type'];
const IMAGES: Record<BottleType, string> = {
  brandy: brandy,
  shochu: shochu,
  whisky: whisky,
};
export const Sake = (props: {
  type: BottleType;
  width?: number | string;
  height?: number | string;
}) => (
  <img
    src={IMAGES[props.type]}
    width={props.width ?? 60}
    height={props.height ?? 100}
    style={{ objectFit: 'contain' }}
  />
);
