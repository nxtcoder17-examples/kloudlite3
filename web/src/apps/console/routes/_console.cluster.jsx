import { useState } from 'react';
import { Link } from '@remix-run/react';
import { PlusFill } from '@jengaicons/react';
import { SubHeader } from '~/components/organisms/sub-header.jsx';
import { Button } from '~/components/atoms/button.jsx';
import { EmptyState } from '~/components/molecule/empty-state.jsx';

const Cluster = () => {
  const [projects, _setProjects] = useState([]);

  return (
    <>
      <SubHeader
        title="Projects"
        actions={
          projects.length !== 0 && (
            <Button
              variant="primary"
              content="Add new"
              prefix={PlusFill}
              href="/new-project"
              LinkComponent={Link}
            />
          )
        }
      />
      {projects.length > 0 && (
        <div className="pt-5 flex flex-col gap-10">
          <div>
            {/* <ResourceList items={[1, 2, 3, 4, 5]} mode={projectListMode} /> */}
          </div>
        </div>
      )}
      {projects.length === 0 && (
        <div className="pt-5">
          <EmptyState
            heading="This is where you’ll manage your projects"
            action={{
              title: 'Create new cluster',
              LinkComponent: Link,
              href: '/cluster',
            }}
          >
            <p>You can create a new project and manage the listed project.</p>
          </EmptyState>
        </div>
      )}
    </>
  );
};

export default Cluster;
