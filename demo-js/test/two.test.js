import { test } from "node:test";
import assert from "node:assert";
import { two } from "../src/two.js";

test("two", () => assert.strictEqual(two(), 2));
