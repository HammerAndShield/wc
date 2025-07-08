import ArgumentParser
import Foundation

@main 
struct Main: AsyncParsableCommand {
    @Flag(name: .customShort("c"), help: "Read the byte count of the file.")
    var countBytes: Bool = false

    @Argument(
        help: "File to be read",
        transform: { URL(filePath: $0) } 
    )
    var fileName: URL

    mutating func run() async {
        if countBytes {
            do {
                let start = Date()
                let fileHandle = try FileHandle(forReadingFrom: fileName)
                defer { try? fileHandle.close() }

                let chunkSize = 4096
                var totalBytes = 0
                while true {
                    let chunk = try fileHandle.read(upToCount: chunkSize)
                    guard let chunk = chunk else {
                        break
                    }
                    totalBytes += chunk.count
                }

                // var totalBytes = 0
                // for try await _ in fileHandle.bytes {
                //     totalBytes += 1
                // }

                let processingTime = Date().timeIntervalSince(start)
                print("\(totalBytes) \(fileName.lastPathComponent)")
                print("processing time: \(processingTime * 1000) ms")
            } catch {
                fatalError("There was an error processing your file: \(error.localizedDescription)")
            }
        } else {
            print("You must specify a command line flag. Check -h for help.")
        }
    }
}
