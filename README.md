# physiology-sim

A 24-hour hackathon project. Simulates blood circulation and the activation of the sympathetic nervous system during exercise and rest.

![Screen Recording 2024-01-28 at 15 04 46](https://github.com/villekuosmanen/physiology-sim/assets/25554034/7c86e0d5-cf4e-4935-87a3-b851e29c6e87)

Simulation is quite thorough: organs and blood vessels are modelled after their real-world counterparts in terms of blood circulation. For response to exercise, the following are modelled:
- intake of oxygen in the lungs, and its consumption through aerobic metabolism in organs
- production of lactic acid during anaerobic metabolism in muscles
- burning / removal of lactic acid under excess oxygen in muscles and the liver
- sensing of blood acidity (co2 + lactic acid) in the autonomous nervous system, and extrection of a neurotransmitter involved in stimulating the sympathetic nervous system, or extrection of a compound to disable the former neurotransmitter.
- effect of the neurotransmitter in different organs, for example in increasing heart and resporation rate and increasing blood flow in blood vessels.

The above are built into the system, and when combined with carefully calibrated constants they create a realistic simulation of human physiology with lots of interesting emergent properties.

<img width="1461" alt="Screenshot 2024-01-28 at 10 38 10" src="https://github.com/villekuosmanen/physiology-sim/assets/25554034/9530ce5f-0d61-4ce2-9361-4fffb603b3aa">
