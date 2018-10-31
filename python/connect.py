#!/usr/bin/env python

# WS client example

import asyncio
import base64
import json
import websockets

ACK = base64.b64encode(b"--ACK--")

async def run():
    async with websockets.connect('ws://localhost:8080/v1/session') as websocket:

        try:
            while True:
                msg = await websocket.recv()
                in_data = json.loads(base64.b64decode(msg))
                
                if 'Uptime' in in_data:
                    await websocket.send(ACK)
                else:
                    print("< {}".format(str(in_data)))
        except websockets.exceptions.ConnectionClosed:
            print("Received close message - Disconnecting...")

asyncio.get_event_loop().run_until_complete(run())