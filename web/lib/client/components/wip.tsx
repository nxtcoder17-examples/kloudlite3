import { Integration } from '@jengaicons/react';
import { EmptyState } from '~/console/components/empty-state';

const Wip = () => {
  return (
    <EmptyState
      heading="Page is under construction"
      image={<Integration size={48} />}
    />
  );
};

export default Wip;
