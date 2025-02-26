# SICSIM

SICSIM is a SIC/XE emulator written in Go. This project replicates the SIC/XE computer architecture.

## Features

- **SIC/XE Simulation:**  
  Emulates the SIC/XE architecture with support for almost all standard instructions.
- **Registers:**  
  Fully implements SIC/XE **registers**, allowing proper execution of machine instructions.
- **Memory Management:**  
  Simulates **memory operations**, allowing programs to store and retrieve data as they would on a real SIC/XE machine.
- **Addressing Modes:**  
  Implements all SIC/XE **addressing modes** to accurately simulate the behavior of instructions.
- **Instruction Formats:**  
  Supports **all instruction formats** used in SIC/XE, ensuring compatibility with different operations.
- **Device Management:**
  - Device `0` → **Standard Input**
  - Device `1` → **Standard Output**
  - Device `2` → **Error Output**
  - Other devices
- **Graphical User Interface (GUI):**  
  Provides a **simple and intuitive GUI** using [Fyne](https://fyne.io/) for easier interaction with the simulator.

## 🖼 Screenshot

![SICSIM Screenshot](assets/simulator.png)

## Limitations

- **No Floating-Point & System Instructions:**  
  It does not support floating-point operations or privileged system instructions.

- **No Re-addressing (Modification Record Handling):**  
  The simulator does not process **modification records** in object files. This means relocatable code requiring address adjustments during loading **will not work correctly**, potentially causing errors when running such programs.

- **No Assembler or Linker Integration:**  
  The simulator does not include an assembler or linker, so source code must be assembled externally before loading into SICSIM.
