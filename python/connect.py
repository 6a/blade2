#!/usr/bin/env python

# WS client example

import asyncio
import base64
import json
import websockets

ACK_BIN = b"--ACK--"
ACK_B64 = base64.b64encode(ACK_BIN)


async def run():
    async with websockets.connect('ws://localhost:8080/v1/session') as websocket:

        try:
            while True:
                msg = await websocket.recv()
                if (msg == ACK_BIN):
                    await websocket.send(ACK_B64)
                else:
                    try:
                        in_data = json.loads(base64.b64decode(msg))
                        print("< {}".format(str(in_data)))
                    except Exception:
                        print("< Unexpected: {}".format(msg))                  
                    
        except websockets.exceptions.ConnectionClosed:
            print("Received close message - Disconnecting...")

asyncio.get_event_loop().run_until_complete(run())