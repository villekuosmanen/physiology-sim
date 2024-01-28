import React, { useState, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import './App.css';

function App() {
  const [socketUrl] = useState('ws://localhost:7766/ws'); // Update with your WebSocket server URL

  const [hoveredElement, setHoveredElement] = useState('');
  const [exerciseLevel, setExerciseLevel] = useState(0); // 0 can be the default value
  const [isFastForward, setIsFastForward] = useState(false);
  const [currentHeartRate, setCurrentHeartRate] = useState(70); // Default heart rate, for example
  const [currentEffort, setCurrentEffort] = useState(1); // Default effort, for example
  const [componentData, setComponentData] = useState({});

  const { sendMessage, lastMessage, readyState } = useWebSocket(socketUrl, {
    onOpen: () => {
      console.log('WebSocket Connected');
    }
  });

  useEffect(() => {
    if (lastMessage !== null) {
      console.log(lastMessage.data)
      const messageData = JSON.parse(lastMessage.data);
      if (messageData.component_name !== undefined) {
        setComponentData(prevData => ({
          ...prevData,
          [messageData.component_name]: messageData
        }));

      } else {
        // general stat
        if (messageData.heart_rate !== undefined) {
          setCurrentHeartRate(messageData.heart_rate)
        }
        if (messageData.effort !== undefined) {
          setCurrentEffort(messageData.effort)
        }
      }
    }
  }, [lastMessage]);



  const handleExerciseChange = (event) => {
    setExerciseLevel(event.target.value);
    sendMessage(JSON.stringify({
      message: 'set_exercise_level',
      level: METValues[event.target.value],
    }));
  };
  
  const toggleFastForward = () => {
    sendMessage(JSON.stringify({
      message: 'toggle_fast_forward'
    }));
    setIsFastForward(!isFastForward);
  };

  const toggleSimulation = () => {
    sendMessage(JSON.stringify({
      message: 'toggle_simulation'
    }));
  };


  function handleMouseEnter(e) {
    setHoveredElement(e.target.id);
  }

  const METValues = {
    0: 0.66,
    1: 1,
    2: 3,
    3: 5.5,
    4: 8,
    5: 10
  };

  const ReverseMETValues = {
    0.66: "Rest",
    1: 'Sitting down',
    3: 'Light Cardio',
    5.5: 'Medium Cardio',
    8: 'Heavy Cardio',
    10: 'Extreme Cardio',
  };

  return (
    <div className="App-header">
      <div className="left">
        <p>Exercise level</p>
        <input type="range" min="0" max="5" value={exerciseLevel} onChange={handleExerciseChange} />
        {['Rest', 'Sitting down', 'Light Cardio', 'Medium Cardio', 'Heavy Cardio', 'Extreme Cardio'][exerciseLevel]}

        <div className='fast-forward'>
          <p>Fast Forward</p>
          <input type="checkbox" checked={isFastForward} onChange={toggleFastForward} />
        </div>

        <button onClick={toggleSimulation}>Reset simulation</button>
      </div>
      <div className="centre">
        <svg viewBox="100 0 400 500" xmlns="http://www.w3.org/2000/svg">
          <rect x="200" y="150" width="100" height="163.611" style={{ fill: 'rgb(216, 216, 216)', stroke: 'rgb(0, 0, 0)' }} transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <rect id="right-breast" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} x="260" y="155" width="35" height="51.949" style={{ stroke: 'rgb(0, 0, 0)', fill: 'rgb(178, 178, 178)' }} transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <rect id="left-breast" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} x="205" y="155" width="35" height="51.749" style={{ stroke: 'rgb(0, 0, 0)', fill: 'rgb(178, 178, 178)' }} transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <rect id="abdomen" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} x="205" y="230" width="90" height="78.484" style={{ stroke: 'rgb(0, 0, 0)', fill: 'rgb(178, 178, 178)' }} transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <ellipse style={{ fill: 'rgb(216, 216, 216)', stroke: 'rgb(0, 0, 0)' }} cx="250" cy="104.585" rx="40.606" ry="40.606" transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <path id="brain" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} d="M 218.049 67.741 C 218.049 94.615 247.141 111.412 270.414 97.974 C 281.216 91.738 287.87 80.213 287.87 67.741" style={{ fill: 'rgb(255, 164, 221)', transformOrigin: '251.5px 85px' }} transform="matrix(-1, 0, 0, -1, 0, -0.000015258789)" />
          <rect id="right-leg" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} x="200" y="300" width="35.973" height="143.246" style={{ fill: 'rgb(216, 216, 216)', stroke: 'rgb(0, 0, 0)', transformOrigin: '200px 380px' }} transform="matrix(0.996195018291, 0.087156012654, -0.087156012654, 0.996195018291, 1.360991642634, 18.602609368772)" />
          <rect id="left-leg" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} x="200" y="300" width="35.973" height="143.246" style={{ fill: 'rgb(216, 216, 216)', stroke: 'rgb(0, 0, 0)', transformOrigin: '200px 380px' }} transform="matrix(0.996195018291, -0.087156012654, 0.087156012654, 0.996195018291, 65.372834425879, 22.396135450998)" />
          <rect id="right-arm" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} x="200" y="310" width="35.973" height="143.246" style={{ fill: 'rgb(216, 216, 216)', stroke: 'rgb(0, 0, 0)', transformOrigin: '220.19px 379.391px' }} transform="matrix(0.996195018291, 0.087156012654, -0.087156012654, 0.996195018291, -45.416666869337, -153.285957302958)" />
          <rect id="left-arm" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} x="200" y="310" width="35.973" height="143.246" style={{ fill: 'rgb(216, 216, 216)', stroke: 'rgb(0, 0, 0)', transformOrigin: '220.19px 379.391px' }} transform="matrix(0.996195018291, -0.087156012654, 0.087156012654, 0.996195018291, 110.998910170508, -153.058561204275)" />
          <path id="left-kidney" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} d="M 265.806 302.665 C 267.226 302.977 268.837 302.984 270.238 302.665 C 271.633 302.348 273.261 301.651 274.194 300.767 C 275.035 299.969 275.506 298.864 275.776 297.76 C 276.058 296.606 276.113 295.143 275.776 293.962 C 275.435 292.768 274.57 291.412 273.719 290.639 C 272.956 289.947 271.649 290.329 271.029 289.373 C 270.123 287.979 269.52 283.218 270.079 281.778 C 270.445 280.838 271.702 281.173 272.295 280.354 C 273.078 279.269 274.045 276.935 273.877 275.29 C 273.705 273.61 272.591 271.417 271.187 270.385 C 269.73 269.313 267.09 268.837 265.174 269.119 C 263.238 269.404 260.993 270.771 259.636 272.126 C 258.361 273.397 257.641 275.002 257.104 276.873 C 256.487 279.023 256.38 281.62 256.471 284.468 C 256.578 287.826 256.943 293.016 258.053 295.861 C 258.915 298.07 260.304 299.62 261.693 300.767 C 262.926 301.785 264.376 302.35 265.806 302.665 Z" style={{ stroke: 'rgb(0, 0, 0)', strokeMiterlimit: 1, fill: 'rgb(158, 172, 51)', transformOrigin: '266.226px 285.969px' }} transform="matrix(-1, 0, 0, -1, 0, 0)" />
          <path id="lungs" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} d="M 230.017 171.467 C 217.263 171.467 209.291 202.384 215.669 227.117 C 218.628 238.596 224.098 245.667 230.017 245.667 C 242.771 245.667 250.742 214.75 244.365 190.017 C 241.406 178.538 235.936 171.467 230.017 171.467 M 269.361 173.057 C 256.607 173.057 248.635 203.974 255.013 228.707 C 257.972 240.186 263.442 247.257 269.361 247.257 C 282.115 247.257 290.086 216.34 283.709 191.607 C 280.75 180.128 275.28 173.057 269.361 173.057 M 244.096 127.112 L 256.606 127.112 L 256.606 219.383 L 244.096 219.383 L 244.096 127.112 Z" stroke="rgb(0, 0, 0)" style={{ fill: 'rgb(203, 142, 255)' }} transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <ellipse id="heart" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} style={{ stroke: 'rgb(0, 0, 0)', fill: 'rgb(222, 0, 0)' }} cx="251.09" cy="214.315" rx="17.394" ry="24.685" transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <path id="liver" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} d="M 256.811 239.149 C 258.583 239.277 263.53 240.942 265.17 243.074 C 266.66 245.011 267.239 248.413 266.801 250.925 C 266.337 253.583 264.601 255.989 262.112 258.545 C 258.694 262.055 251.181 266.737 246.617 269.166 C 242.972 271.106 239.899 272.281 236.83 272.861 C 234.094 273.379 231.79 273.682 229.083 272.861 C 225.923 271.902 221.416 269.022 219.093 266.626 C 217.205 264.681 216.239 262.649 215.423 260.161 C 214.54 257.472 213.841 253.545 214.199 250.925 C 214.509 248.658 216.08 246.195 216.85 245.152 C 217.296 244.548 217.436 244.589 218.073 244.229 C 219.166 243.611 220.45 242.562 223.374 241.919 C 228.827 240.722 248.918 239.397 253.388 239.57 C 253.491 239.545 253.594 239.522 253.695 239.5 C 254.434 239.342 255.25 239.272 256.069 239.293 C 256.286 239.192 256.529 239.128 256.811 239.149 Z" style={{ stroke: 'rgb(0, 0, 0)', strokeMiterlimit: 1, fill: 'rgb(150, 63, 3)' }} transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
          <path id="right-kidney" onMouseEnter={handleMouseEnter} onMouseLeave={() => setHoveredElement('')} d="M 240.216 270.127 C 241.636 269.815 243.247 269.808 244.648 270.127 C 246.043 270.444 247.671 271.141 248.604 272.025 C 249.445 272.823 249.916 273.928 250.186 275.032 C 250.468 276.186 250.523 277.649 250.186 278.83 C 249.845 280.024 248.98 281.38 248.129 282.153 C 247.366 282.845 246.059 282.463 245.439 283.419 C 244.533 284.813 243.93 289.574 244.489 291.014 C 244.855 291.954 246.112 291.619 246.705 292.438 C 247.488 293.523 248.455 295.857 248.287 297.502 C 248.115 299.182 247.001 301.375 245.597 302.407 C 244.14 303.479 241.5 303.955 239.584 303.673 C 237.648 303.388 235.403 302.021 234.046 300.666 C 232.771 299.395 232.051 297.79 231.514 295.919 C 230.897 293.769 230.79 291.172 230.881 288.324 C 230.988 284.966 231.353 279.776 232.463 276.931 C 233.325 274.722 234.714 273.172 236.103 272.025 C 237.336 271.007 238.786 270.442 240.216 270.127 Z" style={{ stroke: 'rgb(0, 0, 0)', strokeMiterlimit: 1, fill: 'rgb(158, 172, 51)' }} transform="matrix(1, 0, 0, 1, 7.105427357601002e-15, 0)" />
        </svg>
      </div>
      <div className="right">
        <p>Current Heart Rate: {currentHeartRate} bpm</p>
        <p>Current Effort (MET): {ReverseMETValues[currentEffort]}</p>
        {hoveredElement && componentData[hoveredElement] ? <div>
          <div>{hoveredElement}</div>
          {componentData[hoveredElement].blood_quantity !== undefined
            ? <div>
              Blood quantity: {(componentData[hoveredElement].blood_quantity).toFixed(2)}
            </div>
            : null}
          {componentData[hoveredElement].has_oxygen_saturation 
            ? <div>
              Blood oxygen: {(componentData[hoveredElement].oxygen_saturation * 100).toFixed(2)}%
            </div>
            : null}
          {componentData[hoveredElement].has_lactic_acid 
            ? <div>
              Lactic Acid: {(componentData[hoveredElement].lactic_acid * 100).toFixed(2)}
            </div>
            : null}
          {componentData[hoveredElement].has_norepinephrine 
            ? <div>
              Norepinephrine: {(componentData[hoveredElement].norepinephrine * 100).toFixed(2)}%
            </div>
            : null}
        </div> : null}
      </div>
    </div>
  );
}

export default App;
