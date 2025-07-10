import ArgumentParser
import Foundation

@main 
struct Main: AsyncParsableCommand {
    @Flag(name: .customShort("c"), help: "Read the byte count of the file.")
    var countBytes: Bool = false

    @Flag(name: .customShort("l"), help: "Read the number of lines in the file.")
    var countLines: Bool = false

    @Flag(name: .customShort("w"), help: "Count the number of words in the file.")
    var countWords: Bool = false

    @Flag(name: .customShort("p"), help: "Print the time it took to process the file.")
    var trackTime: Bool = false

    @Argument(
        help: "File to be read",
        transform: { URL(filePath: $0) } 
    )
    var fileName: URL

    var noArgs: Bool {
        !countBytes && !countLines && !countWords
    }

    mutating func run() async {
        let start: Date = Date()

        do {
            let fileHandle = try FileHandle(forReadingFrom: fileName)
            defer { try? fileHandle.close() }

            // let chunkSize = 4096
            // var totalBytes = 0
            // while true {
            //     let chunk = try fileHandle.read(upToCount: chunkSize)
            //     guard let chunk = chunk else {
            //         break
            //     }
            //     totalBytes += chunk.count
            // }

            var totalLines = 0
            var totalWords = 0
            for try await line in fileHandle.bytes.lines {
                totalLines += 1

                let words = line.components(separatedBy: .whitespacesAndNewlines)
                    .reduce(0) { $0 + ($1.isEmpty ? 0 : 1) }
                totalWords += words
            }

            var res = "  "
            if countLines || noArgs {
                res += "\(totalLines)  "
            }
            if countWords || noArgs {
                res += "\(totalWords)  "
            }
            if countBytes || noArgs {
                let fileAttributes = try FileManager.default.attributesOfItem(atPath: fileName.path())
                let totalBytes = fileAttributes[.size] as? Int64 ?? 0
                res += "\(totalBytes)  "
            }
            res += fileName.lastPathComponent
            print(res)
        } catch {
            fatalError("There was an error processing your file: \(error.localizedDescription)")
        }

        if trackTime {
            let interval = Date().timeIntervalSince(start)
            print("Processing took \(interval * 1000)ms")
        }
    }
}
