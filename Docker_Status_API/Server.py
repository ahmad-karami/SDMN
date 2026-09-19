from http.server import HTTPServer , BaseHTTPRequestHandler
import json
import socket

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.setsockopt(socket.SOL_SOCKET,socket.SO_REUSEADDR,1)

# False while the status is "OK". Each valid POST flips it.
toggle = False
class MyHandler(BaseHTTPRequestHandler): 
    

    def do_GET(self):
        global toggle
        if (not toggle) & (self.path == "/api/v1/status"):
            self.send_response(200)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            response = {"status": "OK"}
            self.wfile.write(json.dumps(response).encode())

        elif (toggle) & (self.path == "/api/v1/status"):
            self.send_response(200)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            response = {"status": "not OK"}
            self.wfile.write(json.dumps(response).encode())
        else:
            self.send_response(400)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            response = {"status": "Bad Request!!"}
            self.wfile.write(json.dumps(response).encode())


    def do_POST(self):
        global toggle
        content_length = int(self.headers['Content-Length'])
        post_data = json.loads(self.rfile.read(content_length))
        if (self.path == "/api/v1/status") & (post_data == { "status": "not OK" }):
            
            self.send_response(201)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            response = {"status": "not OK"}
            self.wfile.write(json.dumps(response).encode())
            toggle = not toggle
      
        else:
            self.send_response(400)
            self.send_header('Content-type', 'application/json')
            self.end_headers()
            response = {"status": "Bad Request!!"}
            self.wfile.write(json.dumps(response).encode())



# Docker usually assigns this address to the first container on its default bridge network.
IP = '172.17.0.2'
PORT = 9999
server = HTTPServer((IP,PORT),MyHandler)
print(f"server is working on{IP}://{PORT}")

server.serve_forever()
server.server_close()
print("server stoped")