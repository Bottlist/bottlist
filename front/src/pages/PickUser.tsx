import { ChoicesLayout } from '../components/ChoicesLayout';

export function PickUser() {
  return (
    <ChoicesLayout
      items={[
        { text: 'お客様', link: '/pick-login' },
        {
          text: 'お店',
        },
      ]}
    />
  );
}
