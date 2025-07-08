// swift-tools-version: 6.1
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "wc",
    platforms: [.macOS(.v14)],
    products: [
        .executable(name: "wc", targets: ["wc"])
    ],
    dependencies: [
        .package(url: "https://github.com/apple/swift-argument-parser", 
                 from: "1.6.1")
    ],
    targets: [
        .executableTarget(
            name: "wc",
            dependencies: [
                .product(name: "ArgumentParser",
                         package: "swift-argument-parser"),
            ]),
    ]
)
