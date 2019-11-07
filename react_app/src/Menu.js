import React from 'react';
import {Link} from 'react-router-dom';

const Menu = () => {
  const liStyle = {
    display: 'inline',
    width: '100px'
  }

  return (
    <div style={{width: '500px', textAlign: 'left'}}>
      <ul style={{display: 'flex'}}>
        <li style={liStyle}><Link to='/todos'>Go to Todo...</Link></li>
      </ul>
    </div>
  )
}

export default Menu