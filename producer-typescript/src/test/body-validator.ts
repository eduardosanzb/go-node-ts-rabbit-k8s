/**
 * main: sender, extra fields.
 * “ts”: Timestamp validation
 * “message”: Json validator “message”
 * DONE: “ipV4Validator”: validates an IP (check the leetcode)
 *  A valid IPv4 address is an IP in the form "x1.x2.x3.x4"
 *  where 0 <= xi <= 255 and xi cannot contain leading zeros.
 *  For example, "192.168.1.1" and "192.168.1.0" are valid IPv4 addresses
 *  but "192.168.01.1", while "192.168.1.00" and "192.168@1.1" are invalid IPv4 addresses.
 *  for the ips we have the next scenarios
 */
import test from 'node:test';
import assert from 'node:assert';
import bodyValidater, { validIpV4, validTimestamp, onlyValidKeys, containsSender, validMessage, } from '../body-validator'

test("ipV4Message validator", async (t:any) => {
  await t.test('should fail because is empty', () => {
    assert.strictEqual(validIpV4(""), false)
  })

  await t.test("should fail because string is incomplete", () => {
    assert.strictEqual(validIpV4("192.168.1."), false)
  })

  await t.test('should fail because we have leading 0s', () => {
    assert.strictEqual(validIpV4("192.168.1.01"), false)
    assert.strictEqual(validIpV4("192.168.1.00"), false)
    assert.strictEqual(validIpV4("192.168.01.1"), false)
  })

  await t.test("should fail because some bits are out of bound ", () => {
    assert.strictEqual(validIpV4("192.168.1.256"), false);
  })

  await t.test("should fail because chars strings", () => {
    assert.strictEqual(validIpV4("123.1y8.1.1"), false);
    //assert.strictEqual(validIpV4("1 3.148.1. 1    "), false);
    //assert.strictEqual(validIpV4(" 192. 168. 1. 1"), false);
  });

  await t.test('should pass because is a valid IP', () => {
    assert.strictEqual(validIpV4("192.168.1.1"), true);
  })

  await t.test('should pass because is a valid IP', () => {
    assert.strictEqual(validIpV4("192.168.1.0"), true);
  })
})

test("timestamp validator", (t:any) => {
  assert.strictEqual(validTimestamp(1530228282), true);
  assert.strictEqual(validTimestamp("1530228282"), true);
  assert.strictEqual(validTimestamp("hola"), false);
  assert.strictEqual(validTimestamp(''), false);
  assert.strictEqual(validTimestamp(0), false);
})

test("contains sender and is a string", async(t:any) => {
  assert.strictEqual(containsSender(0), false);
  assert.strictEqual(containsSender('sdf'), false);
  assert.strictEqual(containsSender([1,2]), false);
  assert.strictEqual(containsSender([]), false);
  assert.strictEqual(containsSender({}), false);
  assert.strictEqual(containsSender({ a: true }), false);
  assert.strictEqual(containsSender({ a: true, sender: null }), false);
  assert.strictEqual(containsSender({ a: true, sender: {} }), false);
  assert.strictEqual(containsSender({ sender: 'eduardo' }), true);
})

test("contains a message key with a JSON inside", async(t:any) => {
  assert.strictEqual(validMessage(0), false);
  assert.strictEqual(validMessage('sdf'), false);
  assert.strictEqual(validMessage([1,2]), false);
  assert.strictEqual(validMessage([]), false);
  assert.strictEqual(validMessage({}), false);
  assert.strictEqual(validMessage({ a: true }), false);
  assert.strictEqual(validMessage({ a: true, message: null }), false);
  assert.strictEqual(validMessage({ a: true, message: {} }), false);
  assert.strictEqual(validMessage({ message: { works: true } }), true);
})

test("only have valid keys", async(t:any) => {
  assert.strictEqual(onlyValidKeys(0), false);
  assert.strictEqual(onlyValidKeys('sdf'), false);
  assert.strictEqual(onlyValidKeys([1,2]), false);
  assert.strictEqual(onlyValidKeys([]), false);
  assert.strictEqual(onlyValidKeys({}), false);
  assert.strictEqual(onlyValidKeys({ a: true }), false);
  assert.strictEqual(onlyValidKeys({ 
            nope: true,
            "ts": "1530228282", 
            "sender": "testy-test-service", 
            "message": { 
                "foo": "bar", 
                "baz": "bang" 
            }, 
            "sent-from-ip": "1.2.3.4", 
            "priority": 2 
          }), false);

  assert.strictEqual(onlyValidKeys({ 
    "ts": "1530228282", 
    "sender": "testy-test-service", 
    "message": { 
      "foo": "bar", 
      "baz": "bang" 
    }, 
    "sent-from-ip": "1.2.3.4", 
    "priority": 2 
  }), true);
  
})


test("main", () => {
  assert.strictEqual(bodyValidater({ 
"ts": "1530228282", 
"sender": "testy-test-service", 
"message": { 
"foo": "bar", 
"baz": "bang" 
}, 
"sent-from-ip": "1.2.3.4", 
"priority": 2 
}

), true);
})
