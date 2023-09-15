import { ChoicesLayout } from '../components/ChoicesLayout';

export function PickUser() {
  return (
    <ChoicesLayout
      items={[
        { text: 'お客様', link: '/users/pick-login' },
        {
          text: 'お店',
          link: '/shops/pick-login',
        },
      ]}
    />
  );
}
