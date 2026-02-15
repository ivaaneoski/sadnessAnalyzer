# MoodMetrics
*A tiny Go app that tells you how much emotional copium you inhaled today.*

---
## What is this?

Ever had a rough day and needed someone to analyze your sadness? Daily Sadness Report™ reads your day in a paragraph, calculates your **sadness dose**, and gives you a **motivational (or brutally honest) comment** based on your sadness levels.

## Requirements

- Go 1.x or higher

## Installation

```bash
git clone https://github.com/ivaaneoski/sadnessAnalyzer.git
cd sadnessAnalyzer
```

## Run

```bash
go run main.go
```

## Build

```bash
go build -o sadness-analyzer
./sadness-analyzer
```

Or use the pre-built executables:
- `sadnessReport.exe`
- `sadnessReport1.1build.exe`

## Usage

Run the program and describe your day:

```bash
go run main.go

## 📸 Demo
```bash

🧠 Welcome to your Sadness Report 🧠
  ______                   __                                                 ______                       __                                         
 /      \                 |  \                                               /      \                     |  \                                        
|  $$$$$$\  ______    ____| $$ _______    ______    _______   _______       |  $$$$$$\ _______    ______  | $$ __    __  ________   ______    ______  
| $$___\$$ |      \  /      $$|       \  /      \  /       \ /       \      | $$__| $$|       \  |      \ | $$|  \  |  \|        \ /      \  /      \ 
 \$$    \   \$$$$$$\|  $$$$$$$| $$$$$$$\|  $$$$$$\|  $$$$$$$|  $$$$$$$      | $$    $$| $$$$$$$\  \$$$$$$\| $$| $$  | $$ \$$$$$$$$|  $$$$$$\|  $$$$$$\
 _\$$$$$$\ /      $$| $$  | $$| $$  | $$| $$    $$ \$$    \  \$$    \       | $$$$$$$$| $$  | $$ /      $$| $$| $$  | $$  /    $$ | $$    $$| $$   \$$
|  \__| $$|  $$$$$$$| $$__| $$| $$  | $$| $$$$$$$$ _\$$$$$$\ _\$$$$$$\      | $$  | $$| $$  | $$|  $$$$$$$| $$| $$__/ $$ /  $$$$_ | $$$$$$$$| $$      
 \$$    $$ \$$    $$ \$$    $$| $$  | $$ \$$     \|       $$|       $$      | $$  | $$| $$  | $$ \$$    $$| $$ \$$    $$|  $$    \ \$$     \| $$      
  \$$$$$$   \$$$$$$$  \$$$$$$$ \$$   \$$  \$$$$$$$ \$$$$$$$  \$$$$$$$        \$$   \$$ \$$   \$$  \$$$$$$$ \$$ _\$$$$$$$ \$$$$$$$$  \$$$$$$$ \$$      
                                                                                                              |  \__| $$                              
                                                                                                               \$$    $$                              
                                                                                                                \$$$$$$                               
Tell me about your day:
Today my cat left me for the neighbor and it rained on my pizza delivery.

Your Sadness Dose Today is: 60%
Moderate copium intake... Stay strong 🫡

Thanks for trusting Daily Sadness Report™. Stay strong, my glorious king! 🫡



```
## How it Works

The app analyzes text input for sadness indicators:
- Calculates sadness percentage based on keywords and sentiment
- Provides tiered responses based on sadness level:
  - Low (0-30%): Minimal copium
  - Medium (30-60%): Moderate copium
  - High (60-100%): Heavy copium dosage

Motivational comments are stored in `copium_comments.go` and selected based on your sadness score.

## Features

- Text-based sentiment analysis
- ASCII art banner for maximum emotional impact
- Motivational (or sarcastic) responses
- Sadness percentage calculator
- No external dependencies

## Implementation Details

- Written in pure Go with standard library
- Custom sadness detection algorithm
- Tiered response system based on emotional state
- Command-line interface for maximum accessibility

---
