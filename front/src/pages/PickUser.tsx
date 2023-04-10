import { Choices } from '../components/Choices';

export function PickUser() {
  return (
    <Choices
      items={[
        { text: 'お客様', link: '/pick-login' },
        {
          text: 'お店',
        },
      ]}
    />
  );
}
