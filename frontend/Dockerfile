FROM mhart/alpine-node

# Set workdir,

WORKDIR /src

# Copy package.json to /src

COPY package.json .

# install dependency
RUN npm i

# Copy current folder to src /src

COPY . .

# Expose Internal Port to 3003
EXPOSE 3003

# Execute Command

CMD ["npm", "start"]
