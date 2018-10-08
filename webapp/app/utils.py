import random
import string
def guid():  # generate user id
    def g4():
        return ''.join(random.choices(string.ascii_lowercase + string.digits, k=4))

    return g4() + g4() + '-' + g4() + '-' + g4() + '-' + g4() + '-' + g4() + g4() + g4()
    pass
