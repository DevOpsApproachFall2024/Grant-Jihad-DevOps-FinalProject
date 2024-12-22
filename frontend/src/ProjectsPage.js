import React, { useEffect, useState } from 'react';

const ProjectsPage = () => {
  const [projects, setProjects] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/projects')
      .then(res => res.json())
      .then(data => setProjects(data));
  }, []);

  return (
    <div>
      <h1>Projects</h1>
      <ul>
        {projects.map((p) => (
          <li key={p.id}>{p.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default ProjectsPage;
