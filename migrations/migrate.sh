#!/bin/bash
migrate -source file://. -database postgres://postgres:password@postgres:5432/postgres?sslmode=disable up