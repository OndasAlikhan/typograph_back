# Use the official Nginx image as the base image
FROM nginx:latest

# Copy your Nginx configuration file to the container
COPY ./nginx.conf /etc/nginx/nginx.conf

# Expose port 80 to the outside world
EXPOSE 80

# Command to run Nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]