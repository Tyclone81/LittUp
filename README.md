# INKA LittUp — Electrical Solutions Website

INKA LittUp is a modern, mobile-responsive business website for **INKA LittUp**, a full-service electrical outfit based in Kisii, Kenya. The site showcases the company's services, portfolio, and contact information, and includes a working contact form, WhatsApp integration, and an image gallery with a lightbox.


## Overview

This project was built as a complete, professional web presence for **INKA LittUp** — an electrical outfit co-founded by Raymond Ogero in 2020. The site highlights the company's values, services, and portfolio while making it easy for potential clients to get in touch.

**Goal:** Build a simple, clean, and functional website that works on both desktop and mobile devices.

**Live URL:** [https://inkalittup.onrender.com/](https://inkalittup.onrender.com/)


## Features

| Feature | Description |
| :--- | :--- |
| **Responsive Design** | Works seamlessly on desktop, tablet, and mobile |
| **Hero Section** | Full-width cover image with key messaging and CTA |
| **About Section** | Company story, values, and stats (projects & experience) |
| **Services Section** | Six core services with hover zoom effects |
| **Gallery** | Project portfolio with images and full-screen lightbox |
| **WhatsApp Button** | Direct chat integration for instant customer contact |
| **Contact Form** | Powered by Formspree — sends messages to company email |
| **Mobile-First** | Optimized for mobile viewing and interactions |


## Tech Stack

| Technology | Purpose |
| :--- | :--- |
| **HTML5** | Page structure and content |
| **CSS3** | Styling, layout, and responsiveness |
| **Go (Golang)** | Simple HTTP server for local development |
| **Formspree** | Contact form backend (email forwarding) |
| **Render** | Hosting and deployment (live site) |

---

## Project Structure
Littup/
├── index.html # Main webpage
├── main.go # Go server for local development
├── README.md # This file
├── static/
│ ├── styles.css # All CSS styles
│ ├── cover.jpg # Hero background image
│ ├── job0.jpg # Gallery images (0–9)
│ ├── job2.jpg
│ ├── ...
│ └── job9.jpg

text

---

## How to Run Locally

### Prerequisites

- **Go** installed on your system ([Download Go](https://go.dev/dl/))

### Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Tyclone81/LittUp.git
   cd LittUp
   
Run the Go server:

bash
go run main.go
Open your browser and visit:

http://localhost:8080
The website should now be running locally.

## Deployment
The live version of this site is deployed on Render and can be accessed at:

https://inkalittup.onrender.com/

Auto-Deploy: Any changes pushed to the main branch on GitHub will automatically trigger a redeployment on Render.

## Key Learnings & Challenges
Challenge	Solution
Hero image cropping on mobile	Used background-size: contain with a dark background color to fill empty space
Lightbox gallery	Used pure CSS :target selector for full-screen image viewing
WhatsApp Integration;  Used WhatsApp deep linking (`https://wa.me/`) with a pre-filled message to enable one-click chat.

## Author
Tyclone81
GitHub: https://github.com/Tyclone81

This project was built as part of the production-grade personal learning project to practice web development, Go, and deployment workflows.


## License
This project is for educational and portfolio purposes. All rights reserved by the author.

