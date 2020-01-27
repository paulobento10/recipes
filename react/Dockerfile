
FROM node

ENV NPM_CONFIG_LOGLEVEL warn
ARG app_env
ENV NODE_ENV $app_env

#RUN mkdir -p /react/react-app
WORKDIR /home/react/react-app
COPY ./react/react-app ./

#CMD ll
# & 
#RUN

COPY ./react/react-app/script.sh /usr/bin/
RUN chmod +x /usr/bin/script.sh
CMD script.sh 

# if dev settings will use create-react start script for hot code relaoding via docker-compose shared volume
# if production setting will build optimized static files and serve using http-server
#CMD if [ ${NODE_ENV} = production ]; \
#	then \
#	npm install -g http-server && \
#	npm run build && \
#	cd build && \
#	hs -p 3000; \
#	else \
#	npm run start; \
#	fi

EXPOSE 3000
