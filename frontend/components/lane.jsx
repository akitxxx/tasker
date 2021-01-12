import React from 'react';
import Card from './card';
import './lane.module.scss';

const Lane = (props) => {

  return (
    <div className="lane col-md-4">
      <div className="laneHeader">
        <h5>Tasks</h5>
      </div>
      {props.taskList.map((task) => {
        return <Card key={task.id} title={task.title} />;
      })}
      <div className="laneFooter">
      </div>
    </div>
  );
};

export default Lane;